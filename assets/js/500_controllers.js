
var bsmWebAppControllers = angular.module('bsmWebAppControllers', []);

bsmWebAppControllers.filter('stringify', function(){ 
    return function(input) {
        return JSON.stringify(input);
    };
});

bsmWebAppControllers.controller('MainCtrl', [
    '$scope', '$sails', 'Page',
    function ($scope, $sails, Page) {
        $scope.Page = Page;
    }
]);

bsmWebAppControllers.controller('WelcomeCtrl', [
    '$scope', '$sails', 'Page',
    function ($scope, $sails, Page) {
        Page.setStyle("frontpage");
        

    }
]);

bsmWebAppControllers.controller('ContactCtrl', [
    '$scope', '$sails', 'Page',
    function ($scope, $sails, Page) {
        Page.setStyle("frontpage");
        

    }
]);

bsmWebAppControllers.controller('TestCtrl', [
    '$scope', '$sails', 'Page',
    function ($scope, $sails, Page) {
        Page.setStyle("test");

    }
]);