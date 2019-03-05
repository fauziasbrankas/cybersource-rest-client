/**
 * CyberSource Flex API
 * Simple PAN tokenization service
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.0
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/PtsV2PayoutsPost201ResponseErrorInformationDetails'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./PtsV2PayoutsPost201ResponseErrorInformationDetails'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.UmsV1UsersGet400Response = factory(root.CyberSource.ApiClient, root.CyberSource.PtsV2PayoutsPost201ResponseErrorInformationDetails);
  }
}(this, function(ApiClient, PtsV2PayoutsPost201ResponseErrorInformationDetails) {
  'use strict';




  /**
   * The UmsV1UsersGet400Response model module.
   * @module model/UmsV1UsersGet400Response
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>UmsV1UsersGet400Response</code>.
   * @alias module:model/UmsV1UsersGet400Response
   * @class
   */
  var exports = function() {
    var _this = this;






  };

  /**
   * Constructs a <code>UmsV1UsersGet400Response</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/UmsV1UsersGet400Response} obj Optional instance to populate.
   * @return {module:model/UmsV1UsersGet400Response} The populated <code>UmsV1UsersGet400Response</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('submitTimeUtc')) {
        obj['submitTimeUtc'] = ApiClient.convertToType(data['submitTimeUtc'], 'String');
      }
      if (data.hasOwnProperty('status')) {
        obj['status'] = ApiClient.convertToType(data['status'], 'String');
      }
      if (data.hasOwnProperty('reason')) {
        obj['reason'] = ApiClient.convertToType(data['reason'], 'String');
      }
      if (data.hasOwnProperty('message')) {
        obj['message'] = ApiClient.convertToType(data['message'], 'String');
      }
      if (data.hasOwnProperty('details')) {
        obj['details'] = ApiClient.convertToType(data['details'], [PtsV2PayoutsPost201ResponseErrorInformationDetails]);
      }
    }
    return obj;
  }

  /**
   * Time of request in UTC. `Format: YYYY-MM-DDThh:mm:ssZ`  Example 2016-08-11T22:47:57Z equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The T separates the date and the time. The Z indicates UTC. 
   * @member {String} submitTimeUtc
   */
  exports.prototype['submitTimeUtc'] = undefined;
  /**
   * The status of the submitted transaction.
   * @member {module:model/UmsV1UsersGet400Response.StatusEnum} status
   */
  exports.prototype['status'] = undefined;
  /**
   * The reason of the status. 
   * @member {module:model/UmsV1UsersGet400Response.ReasonEnum} reason
   */
  exports.prototype['reason'] = undefined;
  /**
   * The detail message related to the status and reason listed above.
   * @member {String} message
   */
  exports.prototype['message'] = undefined;
  /**
   * @member {Array.<module:model/PtsV2PayoutsPost201ResponseErrorInformationDetails>} details
   */
  exports.prototype['details'] = undefined;


  /**
   * Allowed values for the <code>status</code> property.
   * @enum {String}
   * @readonly
   */
  exports.StatusEnum = {
    /**
     * value: "INVALID_REQUEST"
     * @const
     */
    "REQUEST": "INVALID_REQUEST"  };

  /**
   * Allowed values for the <code>reason</code> property.
   * @enum {String}
   * @readonly
   */
  exports.ReasonEnum = {
    /**
     * value: "MISSING_FIELD"
     * @const
     */
    "MISSING_FIELD": "MISSING_FIELD",
    /**
     * value: "INVALID_DATA"
     * @const
     */
    "INVALID_DATA": "INVALID_DATA",
    /**
     * value: "DUPLICATE_REQUEST"
     * @const
     */
    "DUPLICATE_REQUEST": "DUPLICATE_REQUEST",
    /**
     * value: "INVALID_CARD"
     * @const
     */
    "INVALID_CARD": "INVALID_CARD",
    /**
     * value: "INVALID_MERCHANT_CONFIGURATION"
     * @const
     */
    "INVALID_MERCHANT_CONFIGURATION": "INVALID_MERCHANT_CONFIGURATION",
    /**
     * value: "CAPTURE_ALREADY_VOIDED"
     * @const
     */
    "CAPTURE_ALREADY_VOIDED": "CAPTURE_ALREADY_VOIDED",
    /**
     * value: "ACCOUNT_NOT_ALLOWED_CREDIT"
     * @const
     */
    "ACCOUNT_NOT_ALLOWED_CREDIT": "ACCOUNT_NOT_ALLOWED_CREDIT"  };


  return exports;
}));


