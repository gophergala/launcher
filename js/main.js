goog.provide('launcher.homepage');
goog.require('goog.dom');
goog.require('goog.events');

launcher.homepage = function() {
    var button = goog.dom.getElementByClass('button');
};
goog.addSingletonGetter(launcher.homepage);

function initHome() {
    launcher.homepage.getInstance();
}
goog.exportSymbol('initHome', initHome);
