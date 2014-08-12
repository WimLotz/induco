(function () {
    var ProfileController = function ($scope) {
        $scope.message = 'profile page';
    };

    ProfileController.$inject = ['$scope'];
    angular.module('induco').controller('ProfileController', ProfileController);
}());