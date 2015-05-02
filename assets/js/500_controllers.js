
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
        Page.setTitle("Welcome");

    }
]);

bsmWebAppControllers.controller('ContactCtrl', [
    '$scope', '$sails', 'Page',
    function ($scope, $sails, Page) {
        Page.setStyle("frontpage");
        Page.setTitle("Contact Us");

    }
]);

bsmWebAppControllers.controller('TestCtrl', [
    '$scope', '$sails', 'Page',
    function ($scope, $sails, Page) {
        Page.setStyle("test");
        Page.setTitle("Test Page");

    }
]);

bsmWebAppControllers.controller('LessonListCtrl', [
    '$scope', '$sails', 'Page',
    function ($scope, $sails, Page) {
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
            }
        ];
    }
]);