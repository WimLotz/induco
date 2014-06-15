(function () {
    var DashboardController = function ($scope) {
        $scope.message = "dashboard hello every one this is you dashboard"
    };
    DashboardController.$inject = ['$scope'];
    angular.module('induco').controller('DashboardController', DashboardController);
}());
