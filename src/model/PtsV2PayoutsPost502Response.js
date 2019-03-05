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
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.PtsV2PayoutsPost502Response = factory(root.CyberSource.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The PtsV2PayoutsPost502Response model module.
   * @module model/PtsV2PayoutsPost502Response
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>PtsV2PayoutsPost502Response</code>.
   * @alias module:model/PtsV2PayoutsPost502Response
   * @class
   */
  var exports = function() {
    var _this = this;





  };

  /**
   * Constructs a <code>PtsV2PayoutsPost502Response</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/PtsV2PayoutsPost502Response} obj Optional instance to populate.
   * @return {module:model/PtsV2PayoutsPost502Response} The populated <code>PtsV2PayoutsPost502Response</code> instance.
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
   * @member {module:model/PtsV2PayoutsPost502Response.StatusEnum} status
   */
  exports.prototype['status'] = undefined;
  /**
   * The reason of the status. 
   * @member {module:model/PtsV2PayoutsPost502Response.ReasonEnum} reason
   */
  exports.prototype['reason'] = undefined;
  /**
   * The detail message related to the status and reason listed above.
   * @member {String} message
   */
  exports.prototype['message'] = undefined;


  /**
   * Allowed values for the <code>status</code> property.
   * @enum {String}
   * @readonly
   */
  exports.StatusEnum = {
    /**
     * value: "SERVER_ERROR"
     * @const
     */
    "ERROR": "SERVER_ERROR"  };

  /**
   * Allowed values for the <code>reason</code> property.
   * @enum {String}
   * @readonly
   */
  exports.ReasonEnum = {
    /**
     * value: "SYSTEM_ERROR"
     * @const
     */
    "SYSTEM_ERROR": "SYSTEM_ERROR",
    /**
     * value: "SERVER_TIMEOUT"
     * @const
     */
    "SERVER_TIMEOUT": "SERVER_TIMEOUT",
    /**
     * value: "SERVICE_TIMEOUT"
     * @const
     */
    "SERVICE_TIMEOUT": "SERVICE_TIMEOUT",
    /**
     * value: "PROCESSOR_TIMEOUT"
     * @const
     */
    "PROCESSOR_TIMEOUT": "PROCESSOR_TIMEOUT"  };


  return exports;
}));


