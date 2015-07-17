
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
    '$scope', '$routeParams', '$http', 'Page', 'SlideNav',
    function ($scope, $routeParams, $http, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Title Slide");
        $http.get('/api/v1/slide/'+$routeParams.id).
            success(function(data, status, headers, config) {
                $scope.slide = data;
            }).error(function(data, status, headers, config) {
                // handle error
            });
        
        $scope.prev = "/#!/lesson/list";
        $scope.next = "/#!/slide/narration/2";
    }
]);

bsmWebAppControllers.controller('SlideImageCtrl', [
    '$scope', '$routeParams', '$http', 'Page', 'SlideNav',
    function ($scope, $routeParams, $http, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Image Slide");
        $http.get('/api/v1/slide/'+$routeParams.id).
            success(function(data, status, headers, config) {
                $scope.slide = data;
            }).error(function(data, status, headers, config) {
                // handle error
            });

        
        $scope.prev = "/#!/slide/narration/2";
        $scope.next = "/#!/lesson/list";
    }
]);

bsmWebAppControllers.controller('SlideNarrationCtrl', [
    '$scope', '$routeParams', '$http', 'Page', 'SlideNav',
    function ($scope, $routeParams, $http, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Narration Slide");
        $http.get('/api/v1/slide/'+$routeParams.id).
            success(function(data, status, headers, config) {
                $scope.slide = data;
            }).error(function(data, status, headers, config) {
                // handle error
            });

        
        $scope.prev = "/#!/slide/title/0";
        $scope.next = "/#!/slide/image/1";
    }
]);
