
var bsmWebApp = angular.module('bsmWebApp', [
    'ngRoute',
    'ngSails',
    'bsmWebAppControllers',
    'pageFactory'
]);

bsmWebApp.config(
    ['$routeProvider', '$locationProvider',
    function($routeProvider, $locationProvider) {
        $locationProvider.hashPrefix('!');
        $routeProvider.
            when('/', {
                templateUrl: 'parts/welcome.html',
                controller: 'WelcomeCtrl'
            }).
            when('/contact', {
                templateUrl: 'parts/contact.html',
                controller: 'ContactCtrl'
            }).
            when('/test', {
                templateUrl: 'parts/test.html',
                controller: 'TestCtrl'
            }).
            otherwise({
                redirectTo: '/'
            });
    }
]);

