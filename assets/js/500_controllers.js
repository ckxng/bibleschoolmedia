
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

bsmWebAppControllers.controller('LessonListCtrl', [
    '$scope', '$sails', 'Page',
    function ($scope, $sails, Page) {
        Page.setStyle("frontpage");
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