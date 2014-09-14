'use strict';

/* Controllers */

var phonecatControllers = angular.module('phonecatControllers', []);

phonecatControllers.controller('HomeCtrl', ['$scope',
  function($scope) {

  }]);


phonecatControllers.controller('PhoneListCtrl', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {
    $http.get('static/' + $routeParams.phoneId +'.json').success(function(data) {
      $scope.phones = data;
    });

    $http.get('static/' + $routeParams.phoneId +'.json').success(function(data) {
      $scope.user = data;
    });

  }]);


/*
phonecatControllers.controller('PhoneDetailCtrl', ['$scope', '$routeParams',
  function($scope, $routeParams) {
    $scope.phoneId = $routeParams.phoneId;
  }]);
*/
