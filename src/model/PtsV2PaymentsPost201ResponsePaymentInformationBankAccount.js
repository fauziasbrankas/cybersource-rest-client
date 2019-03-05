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
    root.CyberSource.PtsV2PaymentsPost201ResponsePaymentInformationBankAccount = factory(root.CyberSource.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The PtsV2PaymentsPost201ResponsePaymentInformationBankAccount model module.
   * @module model/PtsV2PaymentsPost201ResponsePaymentInformationBankAccount
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>PtsV2PaymentsPost201ResponsePaymentInformationBankAccount</code>.
   * @alias module:model/PtsV2PaymentsPost201ResponsePaymentInformationBankAccount
   * @class
   */
  var exports = function() {
    var _this = this;


  };

  /**
   * Constructs a <code>PtsV2PaymentsPost201ResponsePaymentInformationBankAccount</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/PtsV2PaymentsPost201ResponsePaymentInformationBankAccount} obj Optional instance to populate.
   * @return {module:model/PtsV2PaymentsPost201ResponsePaymentInformationBankAccount} The populated <code>PtsV2PaymentsPost201ResponsePaymentInformationBankAccount</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('correctedAccountNumber')) {
        obj['correctedAccountNumber'] = ApiClient.convertToType(data['correctedAccountNumber'], 'String');
      }
    }
    return obj;
  }

  /**
   * Corrected account number from the ACH verification service, which is described in \"ACH Verification,\" page 25. 
   * @member {String} correctedAccountNumber
   */
  exports.prototype['correctedAccountNumber'] = undefined;



  return exports;
}));


