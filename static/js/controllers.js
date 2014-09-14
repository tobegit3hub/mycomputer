'use strict';

/* The angular application controllers */
var mycomputerControllers = angular.module('mycomputerControllers', []);

/* Todo: add comment */
mycomputerControllers.controller('HomeController', ['$scope',
  function($scope) {
}]);

/* This controller to get user and items from mycomputer beego api */
mycomputerControllers.controller('UserItemsController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {
    $http.get('/api/' + $routeParams.username).success(function(data) {
      $scope.items = data;
    });

/*    $http.get('api/' + $routeParams.username +'/items').success(function(data) {
      $scope.user = data;
    });*/

  }]);
