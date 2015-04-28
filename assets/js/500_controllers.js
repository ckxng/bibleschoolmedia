
var bsmWebAppControllers = angular.module('bsmWebAppControllers', []);

bsmWebAppControllers.filter('stringify', function(){ 
    return function(input) {
        return JSON.stringify(input);
    };
});

bsmWebAppControllers.controller('WelcomeCtrl', [
    '$scope', '$sails',
    function ($scope, $sails) {
        

    }
]);

bsmWebAppControllers.controller('ContactCtrl', [
    '$scope', '$sails',
    function ($scope, $sails) {
        

    }
]);

bsmWebAppControllers.controller('TestCtrl', [
    '$scope', '$sails',
    function ($scope, $sails) {
        

    }
]);