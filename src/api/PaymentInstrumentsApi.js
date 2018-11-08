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
    define(['ApiClient', 'model/Body2', 'model/Body3', 'model/InlineResponse400', 'model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response', 'model/TmsV1PaymentinstrumentsPost201Response'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/Body2'), require('../model/Body3'), require('../model/InlineResponse400'), require('../model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response'), require('../model/TmsV1PaymentinstrumentsPost201Response'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.PaymentInstrumentsApi = factory(root.CyberSource.ApiClient, root.CyberSource.Body2, root.CyberSource.Body3, root.CyberSource.InlineResponse400, root.CyberSource.TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response, root.CyberSource.TmsV1PaymentinstrumentsPost201Response);
  }
}(this, function(ApiClient, Body2, Body3, InlineResponse400, TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response, TmsV1PaymentinstrumentsPost201Response) {
  'use strict';

  /**
   * PaymentInstruments service.
   * @module api/PaymentInstrumentsApi
   * @version 0.0.1
   */

  /**
   * Constructs a new PaymentInstrumentsApi. 
   * @alias module:api/PaymentInstrumentsApi
   * @class
   * @param {module:ApiClient} apiClient Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(configObject, apiClient = undefined) {
    this.apiClient = apiClient || ApiClient.instance;

    this.apiClient.setConfiguration(configObject);


    /**
     * Callback function to receive the result of the tmsV1InstrumentidentifiersTokenIdPaymentinstrumentsGet operation.
     * @callback module:api/PaymentInstrumentsApi~tmsV1InstrumentidentifiersTokenIdPaymentinstrumentsGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieve all Payment Instruments associated with an Instrument Identifier
     * @param {String} profileId The id of a profile containing user specific TMS configuration.
     * @param {String} tokenId The TokenId of an Instrument Identifier.
     * @param {Object} opts Optional parameters
     * @param {String} opts.offset Starting Payment Instrument record in zero-based dataset that should be returned as the first object in the array. Default is 0.
     * @param {String} opts.limit The maximum number of Payment Instruments that can be returned in the array starting from the offset record in zero-based dataset. Default is 20, maximum is 100. (default to 20)
     * @param {module:api/PaymentInstrumentsApi~tmsV1InstrumentidentifiersTokenIdPaymentinstrumentsGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response}
     */
    this.tmsV1InstrumentidentifiersTokenIdPaymentinstrumentsGet = function(profileId, tokenId, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'profileId' is set
      if (profileId === undefined || profileId === null) {
        throw new Error("Missing the required parameter 'profileId' when calling tmsV1InstrumentidentifiersTokenIdPaymentinstrumentsGet");
      }

      // verify the required parameter 'tokenId' is set
      if (tokenId === undefined || tokenId === null) {
        throw new Error("Missing the required parameter 'tokenId' when calling tmsV1InstrumentidentifiersTokenIdPaymentinstrumentsGet");
      }


      var pathParams = {
        'tokenId': tokenId
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit']
      };
      var headerParams = {
        'profile-id': profileId
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/json;charset=utf-8'];
      var returnType = TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response;

      return this.apiClient.callApi(
        '/tms/v1/instrumentidentifiers/{tokenId}/paymentinstruments', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the tmsV1PaymentinstrumentsPost operation.
     * @callback module:api/PaymentInstrumentsApi~tmsV1PaymentinstrumentsPostCallback
     * @param {String} error Error message, if any.
     * @param {module:model/TmsV1PaymentinstrumentsPost201Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create a Payment Instrument
     * @param {String} profileId The id of a profile containing user specific TMS configuration.
     * @param {module:model/Body2} body Please specify the customers payment details for card or bank account.
     * @param {module:api/PaymentInstrumentsApi~tmsV1PaymentinstrumentsPostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/TmsV1PaymentinstrumentsPost201Response}
     */
    this.tmsV1PaymentinstrumentsPost = function(profileId, body, callback) {
      var postBody = body;

      // verify the required parameter 'profileId' is set
      if (profileId === undefined || profileId === null) {
        throw new Error("Missing the required parameter 'profileId' when calling tmsV1PaymentinstrumentsPost");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling tmsV1PaymentinstrumentsPost");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var headerParams = {
        'profile-id': profileId
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/json;charset=utf-8'];
      var returnType = TmsV1PaymentinstrumentsPost201Response;

      return this.apiClient.callApi(
        '/tms/v1/paymentinstruments', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the tmsV1PaymentinstrumentsTokenIdDelete operation.
     * @callback module:api/PaymentInstrumentsApi~tmsV1PaymentinstrumentsTokenIdDeleteCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete a Payment Instrument
     * @param {String} profileId The id of a profile containing user specific TMS configuration.
     * @param {String} tokenId The TokenId of a Payment Instrument.
     * @param {module:api/PaymentInstrumentsApi~tmsV1PaymentinstrumentsTokenIdDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.tmsV1PaymentinstrumentsTokenIdDelete = function(profileId, tokenId, callback) {
      var postBody = null;

      // verify the required parameter 'profileId' is set
      if (profileId === undefined || profileId === null) {
        throw new Error("Missing the required parameter 'profileId' when calling tmsV1PaymentinstrumentsTokenIdDelete");
      }

      // verify the required parameter 'tokenId' is set
      if (tokenId === undefined || tokenId === null) {
        throw new Error("Missing the required parameter 'tokenId' when calling tmsV1PaymentinstrumentsTokenIdDelete");
      }


      var pathParams = {
        'tokenId': tokenId
      };
      var queryParams = {
      };
      var headerParams = {
        'profile-id': profileId
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/json;charset=utf-8'];
      var returnType = null;

      return this.apiClient.callApi(
        '/tms/v1/paymentinstruments/{tokenId}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the tmsV1PaymentinstrumentsTokenIdGet operation.
     * @callback module:api/PaymentInstrumentsApi~tmsV1PaymentinstrumentsTokenIdGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/TmsV1PaymentinstrumentsPost201Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieve a Payment Instrument
     * @param {String} profileId The id of a profile containing user specific TMS configuration.
     * @param {String} tokenId The TokenId of a Payment Instrument.
     * @param {module:api/PaymentInstrumentsApi~tmsV1PaymentinstrumentsTokenIdGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/TmsV1PaymentinstrumentsPost201Response}
     */
    this.tmsV1PaymentinstrumentsTokenIdGet = function(profileId, tokenId, callback) {
      var postBody = null;

      // verify the required parameter 'profileId' is set
      if (profileId === undefined || profileId === null) {
        throw new Error("Missing the required parameter 'profileId' when calling tmsV1PaymentinstrumentsTokenIdGet");
      }

      // verify the required parameter 'tokenId' is set
      if (tokenId === undefined || tokenId === null) {
        throw new Error("Missing the required parameter 'tokenId' when calling tmsV1PaymentinstrumentsTokenIdGet");
      }


      var pathParams = {
        'tokenId': tokenId
      };
      var queryParams = {
      };
      var headerParams = {
        'profile-id': profileId
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/json;charset=utf-8'];
      var returnType = TmsV1PaymentinstrumentsPost201Response;

      return this.apiClient.callApi(
        '/tms/v1/paymentinstruments/{tokenId}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the tmsV1PaymentinstrumentsTokenIdPatch operation.
     * @callback module:api/PaymentInstrumentsApi~tmsV1PaymentinstrumentsTokenIdPatchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/TmsV1PaymentinstrumentsPost201Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update a Payment Instrument
     * @param {String} profileId The id of a profile containing user specific TMS configuration.
     * @param {String} tokenId The TokenId of a Payment Instrument.
     * @param {module:model/Body3} body Please specify the customers payment details.
     * @param {module:api/PaymentInstrumentsApi~tmsV1PaymentinstrumentsTokenIdPatchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/TmsV1PaymentinstrumentsPost201Response}
     */
    this.tmsV1PaymentinstrumentsTokenIdPatch = function(profileId, tokenId, body, callback) {
      var postBody = body;

      // verify the required parameter 'profileId' is set
      if (profileId === undefined || profileId === null) {
        throw new Error("Missing the required parameter 'profileId' when calling tmsV1PaymentinstrumentsTokenIdPatch");
      }

      // verify the required parameter 'tokenId' is set
      if (tokenId === undefined || tokenId === null) {
        throw new Error("Missing the required parameter 'tokenId' when calling tmsV1PaymentinstrumentsTokenIdPatch");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling tmsV1PaymentinstrumentsTokenIdPatch");
      }


      var pathParams = {
        'tokenId': tokenId
      };
      var queryParams = {
      };
      var headerParams = {
        'profile-id': profileId
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/json;charset=utf-8'];
      var returnType = TmsV1PaymentinstrumentsPost201Response;

      return this.apiClient.callApi(
        '/tms/v1/paymentinstruments/{tokenId}', 'PATCH',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
