(function () {
    var SearchController = function ($scope) {
        $scope.message = 'search page';
    };

    SearchController.$inject = ['$scope'];
    angular.module('induco').controller('SearchController', SearchController);
}());
