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
    define(['ApiClient', 'model/PtsV2PaymentsPost201ResponseErrorInformationDetails'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./PtsV2PaymentsPost201ResponseErrorInformationDetails'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.TssV2TransactionsGet200ResponseErrorInformation = factory(root.CyberSource.ApiClient, root.CyberSource.PtsV2PaymentsPost201ResponseErrorInformationDetails);
  }
}(this, function(ApiClient, PtsV2PaymentsPost201ResponseErrorInformationDetails) {
  'use strict';




  /**
   * The TssV2TransactionsGet200ResponseErrorInformation model module.
   * @module model/TssV2TransactionsGet200ResponseErrorInformation
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>TssV2TransactionsGet200ResponseErrorInformation</code>.
   * @alias module:model/TssV2TransactionsGet200ResponseErrorInformation
   * @class
   */
  var exports = function() {
    var _this = this;




  };

  /**
   * Constructs a <code>TssV2TransactionsGet200ResponseErrorInformation</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TssV2TransactionsGet200ResponseErrorInformation} obj Optional instance to populate.
   * @return {module:model/TssV2TransactionsGet200ResponseErrorInformation} The populated <code>TssV2TransactionsGet200ResponseErrorInformation</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('reason')) {
        obj['reason'] = ApiClient.convertToType(data['reason'], 'String');
      }
      if (data.hasOwnProperty('message')) {
        obj['message'] = ApiClient.convertToType(data['message'], 'String');
      }
      if (data.hasOwnProperty('details')) {
        obj['details'] = ApiClient.convertToType(data['details'], [PtsV2PaymentsPost201ResponseErrorInformationDetails]);
      }
    }
    return obj;
  }

  /**
   * The description for this field is not available.
   * @member {String} reason
   */
  exports.prototype['reason'] = undefined;
  /**
   * The description for this field is not available.
   * @member {String} message
   */
  exports.prototype['message'] = undefined;
  /**
   * @member {Array.<module:model/PtsV2PaymentsPost201ResponseErrorInformationDetails>} details
   */
  exports.prototype['details'] = undefined;



  return exports;
}));


