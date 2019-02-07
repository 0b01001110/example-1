import {AccountClient} from './pb/account_grpc_web_pb';
import {OpenAccountRequest, TestStreamingRequest} from './pb/account_pb';


var client = new AccountClient('http://localhost:9900');

var request = new OpenAccountRequest();
request.setAccountId("fake-account-id");
request.setName("fake-account-name");

client.openAccount(request, {}, (err, response) => {
    console.log(response.getAccountId());
});

var streamingRequest = new TestStreamingRequest()
var stream = client.testStreaming(streamingRequest, {});
stream.on('data', function(response) {
    console.log(response.getMessage());
});

stream.on('status', function(status) {
    console.log(status.code);
    console.log(status.details);
    console.log(status.metadata);
});

stream.on('end', function(end) {
    console.log("stream ended");
});
