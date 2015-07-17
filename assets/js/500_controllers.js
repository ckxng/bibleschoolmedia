
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
                
                var len = $scope.lessons.data.length;
                for(var i = 0; i < len; i++) {
                    $scope.lessons.data[i].firstSlide = bsmWebAppControllers.slideToUrl(
                        $scope.lessons.data[i].id,
                        0,
                        "slide.TitleSlide");
                }
            }).error(function(data, status, headers, config) {
                // handle error
            });
    }
]);

bsmWebAppControllers.slideToUrl = function(lessonid, seqno, slidetype) {
    var type = "";
    switch(slidetype) {
        case "slide.TitleSlide":
            type = "title";
            break;
        case "slide.ImageSlide":
            type = "image";
            break;
        case "slide.NarrationSlide":
            type = "narration";
            break;
    }
    if(type=="") {
        return null;
    }
    
    return "/#!/lesson/"+lessonid+"/slide/"+seqno+"/"+type;
}

bsmWebAppControllers.updateLessonNav = function($scope, curseqno) {
    var len = $scope.lesson.data.slides.length;
    if(curseqno == 0) {
        $scope.prev = "/#!/lesson/list";
    } else {
        $scope.prev = bsmWebAppControllers.slideToUrl(
            $scope.lesson.data.id,
            parseInt(curseqno)-1,
            $scope.lesson.data.slides[parseInt(curseqno)-1].myType);
    }
    if(curseqno == len-1) {
        $scope.next = "/#!/lesson/list";
    } else {
        $scope.next = bsmWebAppControllers.slideToUrl(
            $scope.lesson.data.id,
            parseInt(curseqno)+1,
            $scope.lesson.data.slides[parseInt(curseqno)+1].myType);
    }
}

bsmWebAppControllers.controller('SlideTitleCtrl', [
    '$scope', '$routeParams', '$http', '$log', 'Page', 'SlideNav',
    function ($scope, $routeParams, $http, $log, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Title Slide");
        
        var lesson = null;
        $http.get('/api/v1/lesson/'+$routeParams.lessonid).
            success(function(data, status, headers, config) {
                $scope.lesson = data;
                
                $http.get('/api/v1/slide/'+$scope.lesson.data.slides[$routeParams.seqno].id).
                    success(function(data, status, headers, config) {
                        $scope.slide = data;
                        
                        bsmWebAppControllers.updateLessonNav($scope, $routeParams.seqno);
                    }).error(function(data, status, headers, config) {
                        // handle error
                    });
                
        
            }).error(function(data, status, headers, config) {
                // handle error
            });
        
    }
]);

bsmWebAppControllers.controller('SlideImageCtrl', [
    '$scope', '$routeParams', '$http', 'Page', 'SlideNav',
    function ($scope, $routeParams, $http, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Image Slide");
        
        var lesson = null;
        $http.get('/api/v1/lesson/'+$routeParams.lessonid).
            success(function(data, status, headers, config) {
                $scope.lesson = data;
                
                $http.get('/api/v1/slide/'+$scope.lesson.data.slides[$routeParams.seqno].id).
                    success(function(data, status, headers, config) {
                        $scope.slide = data;
                        
                        bsmWebAppControllers.updateLessonNav($scope, $routeParams.seqno);
                    }).error(function(data, status, headers, config) {
                        // handle error
                    });
                
        
            }).error(function(data, status, headers, config) {
                // handle error
            });
    }
]);

bsmWebAppControllers.controller('SlideNarrationCtrl', [
    '$scope', '$routeParams', '$http', 'Page', 'SlideNav',
    function ($scope, $routeParams, $http, Page, SlideNav) {
        Page.setStyle("slide");
        Page.setTitle("Sample Narration Slide");
        
        var lesson = null;
        $http.get('/api/v1/lesson/'+$routeParams.lessonid).
            success(function(data, status, headers, config) {
                $scope.lesson = data;
                
                $http.get('/api/v1/slide/'+$scope.lesson.data.slides[$routeParams.seqno].id).
                    success(function(data, status, headers, config) {
                        $scope.slide = data;
                        
                        bsmWebAppControllers.updateLessonNav($scope, $routeParams.seqno);
                    }).error(function(data, status, headers, config) {
                        // handle error
                    });
                
        
            }).error(function(data, status, headers, config) {
                // handle error
            });
    }
]);
