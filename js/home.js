goog.provide('launcher.homepage');
goog.require('goog.dom');
goog.require('goog.events');
goog.require('goog.net.WebSocket');

launcher.homepage = function() {
    var ws = new goog.net.WebSocket();
    this.ws = ws;
    goog.events.listen(ws, goog.net.WebSocket.EventType.MESSAGE,
			     this.onMessage, undefined, this);
    try {
	ws.open('ws://localhost:8080/ws');
    } catch (e) {
	console.log(e);
    }
    this.addButtonListener();
};
goog.addSingletonGetter(launcher.homepage);

launcher.homepage.prototype.addButtonListener = function() {
    var button = goog.dom.getElementByClass('button');
    goog.events.listen(button, goog.events.EventType.CLICK,
			     this.buttonHandler, undefined, this);
};

launcher.homepage.prototype.buttonHandler = function() {
    this.ws.send("1");
};

launcher.homepage.prototype.onMessage = function(e) {
    var output = goog.dom.getElementByClass('output');
    debugger;
    output.innerHTML += e.getMessage();
};

// function initHome() {
//     launcher.homepage.getInstance();
// }
// goog.exportSymbol('initHome', initHome);
launcher.homepage.getInstance();
