
'use strict';

/* The angular application controllers */
var mycomputerControllers = angular.module('mycomputerControllers', []);

/* This controller to get comment from beego api */
mycomputerControllers.controller('HomeController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {
    /* Get the comment objects */
    $http.get('/api/comment').success(function(data) {
      /* If the data is empty string, don't return objects */
      if(typeof data.Id == "undefined") {
        $scope.comments = null;
      } else {
        $scope.comments = data;
      }
    });
}]);

/* This controller to get user and items from beego api */
mycomputerControllers.controller('UserItemsController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {
    /* Get the user object */
    $http.get('/api/' + $routeParams.username).success(function(data) {
      /* If the data is empty string, don't return objects */
      if(typeof data.Name == "undefined") {
        $scope.user = null;
      } else {
        $scope.user = data;
      }
    });

    /* Get the user item objects */
    $http.get('/api/' + $routeParams.username +'/items').success(function(data) {
      /* If the data is empty string, don't return objects */
      if (typeof data[0].Username == "undefined") {
        $scope.items = null;
      } else {
        $scope.items = data;
      }
    });

    /* Determine whether pop up the form to add item or not */
    $scope.adding = false;
  }
]);
