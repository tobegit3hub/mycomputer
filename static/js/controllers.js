'use strict';

/* Controllers */

var mycomputerControllers = angular.module('mycomputerControllers', []);

mycomputerControllers.controller('HomeController', ['$scope',
  function($scope) {

  }]);


mycomputerControllers.controller('UserItemsController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {
//    $http.get('/api/' + $routeParams.username).success(function(data) {
    $http.get('/static/' + $routeParams.username + ".json").success(function(data) {
      $scope.phones = data;
    });

/*    $http.get('static/' + $routeParams.username +'.json').success(function(data) {
      $scope.user = data;
    });
*/
  }]);
