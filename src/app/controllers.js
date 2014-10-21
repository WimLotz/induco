(function () {

    var TagsController = function ($scope, tagsStorage) {
//        $scope.tags = tagsStorage.getTags();
        $scope.tags = [];

        $scope.addTag = function () {
            $scope.tags.push({id: $scope.tagsCounter, name: $scope.tagName});
            $scope.tagsCounter++;
        };

        $scope.remove = function (tagId) {
            var indexOfTagToRemove = 0;
            $scope.tags.forEach(function (tag) {
                if (tag.id === tagId) {
                    $scope.tags.splice(indexOfTagToRemove, 1);
                    return;
                }
                indexOfTagToRemove++
            });
        };
    };

    var DashboardController = function ($scope) {
        $scope.message = "dashboard page";
    };

    var HomeController = function ($scope, $modal, inducoApi, $location) {
        $scope.user = {};

        $scope.login = function () {
            inducoApi.login($scope.user).success(function () {
                $location.path('profile');
                //todo comes in here even if 500 for no user
            });
        };

        $scope.open = function () {
            $modal.open({
                templateUrl: 'app/modals/create_user.html',
                controller: CreateUserController
            });
        };
    };

    var CreateUserController = function ($scope, $modalInstance, inducoApi) {
        $scope.user = {};

        $scope.ok = function () {
            inducoApi.saveUser($scope.user);
            $modalInstance.close();
        };

        $scope.cancel = function () {
            $modalInstance.dismiss('cancel');
        };
    };

    var IndexController = function ($scope, $location) {
        $scope.isHomePage = function () {
            return $location.path() === '/';
        };
    };

    var NavigationBarController = function ($scope, inducoApi) {
        $scope.signOut = function () {
            inducoApi.signOut();
        };
    };

    var SearchController = function ($scope) {
        $scope.message = 'search page';
    };

    var ProfileController = function ($scope, inducoApi, tagsStorage, $modal) {
        $scope.profiles = [];
        $scope.tagsCounter = 0;

        inducoApi.fetchUserProfiles().success(function (data) {
            data.forEach(function (profile) {
                if (profile.isCompany) {
                    profile.heading = profile.companyName
                } else {
                    profile.heading = profile.firstName + ' ' + profile.surname
                }
                $scope.profiles.push(profile);
            });
        });

        $scope.createPersonProfile = function () {
            $modal.open({
                templateUrl: 'app/modals/create_person_profile.html',
                controller: CreatePersonProfileController
            });
        };

        $scope.createCompanyProfile = function () {
            $modal.open({
                templateUrl: 'app/modals/create_company_profile.html',
                controller: CreateCompanyProfileController
            });
        };
    };

    var CreatePersonProfileController = function ($scope, $modalInstance, inducoApi) {
        $scope.personProfile = {};

        $scope.submitPersonProfileForm = function () {
            $scope.personProfile.IsCompany = false;
            inducoApi.saveProfile($scope.personProfile);
            $modalInstance.close();
        };

        $scope.cancel = function () {
            $modalInstance.dismiss('cancel');
        };
    };

    var CreateCompanyProfileController = function ($scope, $modalInstance, inducoApi) {
        $scope.companyProfile = {};

        $scope.submitCompanyProfileForm = function () {
            $scope.companyProfile.IsCompany = true;
            inducoApi.saveProfile($scope.companyProfile);
            $modalInstance.close();
        };

        $scope.cancel = function () {
            $modalInstance.dismiss('cancel');
        };
    };

    ProfileController.$inject = ['$scope', 'inducoApi', 'tagsStorage', '$modal'];
    DashboardController.$inject = ['$scope'];
    HomeController.$inject = ['$scope', '$modal', 'inducoApi', '$location'];
    NavigationBarController.$inject = ['$scope', 'inducoApi'];
    SearchController.$inject = ['$scope'];
    TagsController.$inject = ['$scope', 'tagsStorage'];
    CreateUserController.$inject = ['$scope', '$modalInstance', 'inducoApi'];
    CreatePersonProfileController.$inject = ['$scope', '$modalInstance', 'inducoApi'];
    CreateCompanyProfileController.$inject = ['$scope', '$modalInstance', 'inducoApi'];
    IndexController.$inject = ['$scope', '$location'];

    angular.module("controllers", [])
        .controller('DashboardController', DashboardController)
        .controller('HomeController', HomeController)
        .controller('NavigationBarController', NavigationBarController)
        .controller('SearchController', SearchController)
        .controller('TagsController', TagsController)
        .controller('CreateUserController', CreateUserController)
        .controller('CreatePersonProfileController', CreatePersonProfileController)
        .controller('CreateCompanyProfileController', CreateCompanyProfileController)
        .controller('IndexController', IndexController)
        .controller('ProfileController', ProfileController);
})();