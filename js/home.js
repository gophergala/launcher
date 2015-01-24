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
    this.addButtonListeners();
};
goog.addSingletonGetter(launcher.homepage);

launcher.homepage.prototype.addButtonListeners = function() {
    var button = goog.dom.getElementByClass('button');
    goog.events.listen(button, goog.events.EventType.CLICK,
			     this.buttonHandler, undefined, this);
    var clear = goog.dom.getElementByClass('clear'); 
    goog.events.listen(clear, goog.events.EventType.CLICK,
			     this.clearHandler, undefined, this);   
};

launcher.homepage.prototype.getWs = function() {
    return this.ws;
};

launcher.homepage.prototype.buttonHandler = function() {
    this.getWs().send("1");
};

launcher.homepage.prototype.clearHandler = function() {
    var output = goog.dom.getElementByClass('output');
    goog.dom.setTextContent(output, '');
};

launcher.homepage.prototype.onMessage = function(e) {
    var output = goog.dom.getElementByClass('output');
    output.innerHTML += e.message;
    output.scrollTop = output.scrollHeight;
};

// function initHome() {
//     launcher.homepage.getInstance();
// }
// goog.exportSymbol('initHome', initHome);
launcher.homepage.getInstance();
