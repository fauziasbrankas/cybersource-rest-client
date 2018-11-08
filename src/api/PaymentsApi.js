/**
 * CyberSource Flex API
 * Simple PAN tokenization service
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/CreatePaymentRequest', 'model/PtsV2PaymentsPost201Response', 'model/PtsV2PaymentsPost400Response', 'model/PtsV2PaymentsPost502Response'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/CreatePaymentRequest'), require('../model/PtsV2PaymentsPost201Response'), require('../model/PtsV2PaymentsPost400Response'), require('../model/PtsV2PaymentsPost502Response'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.PaymentsApi = factory(root.CyberSource.ApiClient, root.CyberSource.CreatePaymentRequest, root.CyberSource.PtsV2PaymentsPost201Response, root.CyberSource.PtsV2PaymentsPost400Response, root.CyberSource.PtsV2PaymentsPost502Response);
  }
}(this, function(ApiClient, CreatePaymentRequest, PtsV2PaymentsPost201Response, PtsV2PaymentsPost400Response, PtsV2PaymentsPost502Response) {
  'use strict';

  /**
   * Payments service.
   * @module api/PaymentsApi
   * @version 0.0.1
   */

  /**
   * Constructs a new PaymentsApi. 
   * @alias module:api/PaymentsApi
   * @class
   * @param {module:ApiClient} apiClient Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(configObject, apiClient = undefined) {
    this.apiClient = apiClient || ApiClient.instance;

    this.apiClient.setConfiguration(configObject);


    /**
     * Callback function to receive the result of the createPayment operation.
     * @callback module:api/PaymentsApi~createPaymentCallback
     * @param {String} error Error message, if any.
     * @param {module:model/PtsV2PaymentsPost201Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Process a Payment
     * Authorize the payment for the transaction. 
     * @param {module:model/CreatePaymentRequest} createPaymentRequest 
     * @param {module:api/PaymentsApi~createPaymentCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/PtsV2PaymentsPost201Response}
     */
    this.createPayment = function(createPaymentRequest, callback) {
      var postBody = createPaymentRequest;

      // verify the required parameter 'createPaymentRequest' is set
      if (createPaymentRequest === undefined || createPaymentRequest === null) {
        throw new Error("Missing the required parameter 'createPaymentRequest' when calling createPayment");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/hal+json;charset=utf-8'];
      var returnType = PtsV2PaymentsPost201Response;

      return this.apiClient.callApi(
        '/pts/v2/payments/', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
