const grpc = require("@grpc/grpc-js");
const mqtt = require('mqtt') 
var protoLoader = require("@grpc/proto-loader");

const PROTO_PATH = "./news.proto";
const mqttClient = mqtt.connect("mqtt://127.0.0.1") 
const topicName = 'grpc_message' 

const options = {
	keepCase: true,
	longs: String,
	enums: String,
	defaults: true,
	oneofs: true,
};

var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);

const Traitement = grpc.loadPackageDefinition(packageDefinition).Traitement;

const grpcClient = new Traitement(
	"localhost:50051",
	grpc.credentials.createInsecure()
);



mqttClient.on('connect', () => { 
	mqttClient.subscribe(topicName, (err, granted) => { 
		if(err)
			console.log(err, 'err'); 
		
      console.log(granted, 'granted') 
  }) 
}) 

mqttClient.on('message', (topic, mqttMessage, packet) => { 
	if(topic === topicName) { 
		const payload = {"message": mqttMessage.toString()};

		console.log("--------");
		console.log(" - Mqtt message receive : " + mqttMessage); 
		console.log(" - Grpc call :");
		console.log(payload);
		
		grpcClient.Action(payload, (error, result) => {
			if(error) 
			console.log(error, 'error'); 
			
			console.log(" - Grpc result :");
        	console.log(result);
    	})
  	} 
}) 
