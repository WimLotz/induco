(function () {
    var DashboardController = function ($scope) {
        $scope.message = "dashboard page";
    };

    DashboardController.$inject = ['$scope'];
    angular.module('induco').controller('DashboardController', DashboardController);
}());