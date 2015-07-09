var slideNavFactory = angular.module('slideNavFactory', []);

slideNavFactory.factory('SlideNav', function() {
   var prev = '';
   var next = '';
   return {
     prev: function() { return prev; },
     setPrev: function(x) { prev = x },
     
     next: function() { return next; },
     setNext: function(x) { next = x }
   };
});
