(function () {
    var ProfileController = function ($scope) {
        $scope.message = 'another useless message for testing';
    };
    ProfileController.$inject = ['$scope'];
    angular.module('induco').controller('ProfileController', ProfileController);
}());