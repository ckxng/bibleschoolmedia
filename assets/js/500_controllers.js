
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
    '$scope', 'Page',
    function ($scope, Page) {
        Page.setStyle("frontpage");
        Page.setTitle("Lesson List");
        $scope.lessons = [
            {
                'id': 1,
                'name': 'God Made the World'
            },
            {
                'id': 2,
                'name': 'Jesus Loves the Little Children'
            },
            {
                'id': 3,
                'name': 'Sample Title 3'
            },
            {
                'id': 4,
                'name': 'Sample Title 4'
            },
            {
                'id': 5,
                'name': 'Sample Title 5'
            },
            {
                'id': 6,
                'name': 'Sample Title 6'
            },
            {
                'id': 7,
                'name': 'Sample Title 7'
            },
            {
                'id': 8,
                'name': 'Sample Title 8'
            },
            {
                'id': 9,
                'name': 'Sample Title 9'
            }
        ];
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
