var pageFactory = angular.module('pageFactory', []);

pageFactory.factory('Page', function() {
   var title = '';
   var style = '';
   return {
     title: function() { return title; },
     setTitle: function(x) { title = x },
     
     style: function() { return style; },
     setStyle: function(x) { style = x }
   };
});