
var bsmWebAppControllers = angular.module('bsmWebAppControllers', []);

bsmWebAppControllers.filter('stringify', function(){ 
    return function(input) {
        return JSON.stringify(input);
    };
});

bsmWebAppControllers.controller('MainCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        $scope.Page = Page;
    }
]);

bsmWebAppControllers.controller('WelcomeCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setStyle("frontpage");
        Page.setTitle("Welcome");

    }
]);

bsmWebAppControllers.controller('ContactCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setStyle("frontpage");
        Page.setTitle("Contact Us");

    }
]);

bsmWebAppControllers.controller('TestCtrl', [
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setStyle("test");
        Page.setTitle("Test Page");

    }
]);

bsmWebAppControllers.controller('LessonListCtrl', [
    '$scope', '$http', 'Page',
    function ($scope, $http, Page) {
        Page.setStyle("frontpage");
        Page.setTitle("Lesson List");
        $http.get('/api/v1/lesson').
            success(function(data, status, headers, config) {
                $scope.lessons = data;
            }).error(function(data, status, headers, config) {
                // handle error
            });
    }
]);

bsmWebAppControllers.controller('SlideTitleCtrl', [
    '$scope', '$routeParams', 'Page', 'SlideNav',
    function ($scope, $routeParams, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Title Slide");
        
        $scope.prev = "/#!/lesson/list";
        $scope.next = "/#!/slide/narration/1";

        $scope.id = $routeParams.id;
    }
]);

bsmWebAppControllers.controller('SlideImageCtrl', [
    '$scope', '$routeParams', 'Page', 'SlideNav',
    function ($scope, $routeParams, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Image Slide");
        
        $scope.prev = "/#!/slide/narration/1";
        $scope.next = "/#!/lesson/list";

        $scope.id = $routeParams.id;
    }
]);

bsmWebAppControllers.controller('SlideNarrationCtrl', [
    '$scope', '$routeParams', 'Page', 'SlideNav',
    function ($scope, $routeParams, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Narration Slide");
        
        $scope.prev = "/#!/slide/title/1";
        $scope.next = "/#!/slide/image/1";

        $scope.id = $routeParams.id;
    }
]);
