
var bsmWebApp = angular.module('bsmWebApp', [
    'ngRoute',
    'bsmWebAppControllers',
    'pageFactory',
    'slideNavFactory'
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
            when('/lesson/list', {
                templateUrl: 'parts/lesson-list.html',
                controller: 'LessonListCtrl'
            }).
            when('/slide/title/:id', {
                templateUrl: 'parts/slide-title.html',
                controller: 'SlideTitleCtrl'
            }).
            when('/slide/image/:id', {
                templateUrl: 'parts/slide-image.html',
                controller: 'SlideImageCtrl'
            }).
            when('/slide/narration/:id', {
                templateUrl: 'parts/slide-narration.html',
                controller: 'SlideNarrationCtrl'
            }).
            otherwise({
                redirectTo: '/'
            });
    }
]);

