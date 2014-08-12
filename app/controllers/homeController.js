(function () {
    var HomeController = function ($scope, $http, $window, $location) {
        $scope.isLoggedIn = false;

//        $http({method: 'GET', url: 'http://localhost:8080/isLoggedIn'})
//            .success(function (data) {
//                $scope.isLoggedIn = data.isLoggedIn;
//
//                if($scope.isLoggedIn){
//                    $location.path('dashboard');
//                }
//            })
//            .error(function (data) {
//                console.log('An error has occurred: ' + data);
//            });
    };

    HomeController.$inject = ['$scope', '$http', '$window', '$location'];
    angular.module('induco').controller('HomeController', HomeController);
}());