import {k8sPatch} from '../../module/k8s';

angular.module('bridge.page')
.controller('ConfigureLabelsCtrl', function(_, $scope, $uibModalInstance, $controller, $rootScope, resource, kind) {
  'use strict';

  $scope.resource = resource;
  $scope.kind = kind;
  $scope.fields = {
    labels: angular.copy(resource.metadata.labels),
  };

  $scope.iconProps = {kind: kind.id};

  $scope.save = function() {
    var patch = [{ op: 'replace', path: '/metadata/labels', value: $scope.fields.labels }];
    $scope.requestPromise = k8sPatch(kind, resource, patch);
    $scope.requestPromise.then(function(result) {
      $uibModalInstance.close(result);
    });
  };

  $scope.cancel = function() {
    $uibModalInstance.dismiss('cancel');
  };

})
.controller('ConfigureLabelsFormCtrl', function($scope) {
  'use strict';
  $scope.submit = $scope.save;
});
