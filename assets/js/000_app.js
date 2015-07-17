
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
            when('/lesson/:lessonid/slide/:seqno/title', {
                templateUrl: 'parts/slide-title.html',
                controller: 'SlideTitleCtrl'
            }).
            when('/lesson/:lessonid/slide/:seqno/image', {
                templateUrl: 'parts/slide-image.html',
                controller: 'SlideImageCtrl'
            }).
            when('/lesson/:lessonid/slide/:seqno/narration', {
                templateUrl: 'parts/slide-narration.html',
                controller: 'SlideNarrationCtrl'
            }).
            otherwise({
                redirectTo: '/'
            });
    }
]);

