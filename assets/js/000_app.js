
var bsmWebApp = angular.module('bsmWebApp', [
    'ngRoute',
    'ngSails',
    'bsmWebAppControllers'
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
            otherwise({
                redirectTo: '/'
            });
    }
]);

