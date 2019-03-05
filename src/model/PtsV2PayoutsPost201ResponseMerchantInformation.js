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
    define(['ApiClient', 'model/PtsV2PayoutsPost201ResponseMerchantInformationMerchantDescriptor'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./PtsV2PayoutsPost201ResponseMerchantInformationMerchantDescriptor'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.PtsV2PayoutsPost201ResponseMerchantInformation = factory(root.CyberSource.ApiClient, root.CyberSource.PtsV2PayoutsPost201ResponseMerchantInformationMerchantDescriptor);
  }
}(this, function(ApiClient, PtsV2PayoutsPost201ResponseMerchantInformationMerchantDescriptor) {
  'use strict';




  /**
   * The PtsV2PayoutsPost201ResponseMerchantInformation model module.
   * @module model/PtsV2PayoutsPost201ResponseMerchantInformation
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>PtsV2PayoutsPost201ResponseMerchantInformation</code>.
   * @alias module:model/PtsV2PayoutsPost201ResponseMerchantInformation
   * @class
   */
  var exports = function() {
    var _this = this;


  };

  /**
   * Constructs a <code>PtsV2PayoutsPost201ResponseMerchantInformation</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/PtsV2PayoutsPost201ResponseMerchantInformation} obj Optional instance to populate.
   * @return {module:model/PtsV2PayoutsPost201ResponseMerchantInformation} The populated <code>PtsV2PayoutsPost201ResponseMerchantInformation</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('merchantDescriptor')) {
        obj['merchantDescriptor'] = PtsV2PayoutsPost201ResponseMerchantInformationMerchantDescriptor.constructFromObject(data['merchantDescriptor']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/PtsV2PayoutsPost201ResponseMerchantInformationMerchantDescriptor} merchantDescriptor
   */
  exports.prototype['merchantDescriptor'] = undefined;



  return exports;
}));


