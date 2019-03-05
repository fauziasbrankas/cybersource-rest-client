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
    define(['ApiClient', 'model/Ptsv2paymentsIssuerInformation'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Ptsv2paymentsIssuerInformation'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.Ptsv2paymentsidreversalsProcessingInformation = factory(root.CyberSource.ApiClient, root.CyberSource.Ptsv2paymentsIssuerInformation);
  }
}(this, function(ApiClient, Ptsv2paymentsIssuerInformation) {
  'use strict';




  /**
   * The Ptsv2paymentsidreversalsProcessingInformation model module.
   * @module model/Ptsv2paymentsidreversalsProcessingInformation
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>Ptsv2paymentsidreversalsProcessingInformation</code>.
   * @alias module:model/Ptsv2paymentsidreversalsProcessingInformation
   * @class
   */
  var exports = function() {
    var _this = this;







  };

  /**
   * Constructs a <code>Ptsv2paymentsidreversalsProcessingInformation</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Ptsv2paymentsidreversalsProcessingInformation} obj Optional instance to populate.
   * @return {module:model/Ptsv2paymentsidreversalsProcessingInformation} The populated <code>Ptsv2paymentsidreversalsProcessingInformation</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('paymentSolution')) {
        obj['paymentSolution'] = ApiClient.convertToType(data['paymentSolution'], 'String');
      }
      if (data.hasOwnProperty('reconciliationId')) {
        obj['reconciliationId'] = ApiClient.convertToType(data['reconciliationId'], 'String');
      }
      if (data.hasOwnProperty('linkId')) {
        obj['linkId'] = ApiClient.convertToType(data['linkId'], 'String');
      }
      if (data.hasOwnProperty('reportGroup')) {
        obj['reportGroup'] = ApiClient.convertToType(data['reportGroup'], 'String');
      }
      if (data.hasOwnProperty('visaCheckoutId')) {
        obj['visaCheckoutId'] = ApiClient.convertToType(data['visaCheckoutId'], 'String');
      }
      if (data.hasOwnProperty('issuer')) {
        obj['issuer'] = Ptsv2paymentsIssuerInformation.constructFromObject(data['issuer']);
      }
    }
    return obj;
  }

  /**
   * Type of digital payment solution for the transaction. Possible Values:   - **visacheckout**: Visa Checkout. This value is required for Visa Checkout transactions. See Visa Checkout Using the SCMP API.  - **005**: Masterpass. This value is required for Masterpass transactions on OmniPay Direct. See \"Masterpass,\" page 153. 
   * @member {String} paymentSolution
   */
  exports.prototype['paymentSolution'] = undefined;
  /**
   * Please check with Cybersource customer support to see if your merchant account is configured correctly so you can include this field in your request. * For Payouts: max length for FDCCompass is String (22). 
   * @member {String} reconciliationId
   */
  exports.prototype['reconciliationId'] = undefined;
  /**
   * Value that links the current authorization request to the original authorization request. Set this value to the ID that was returned in the reply message from the original authorization request.  This value is used for:   - Partial authorizations: See \"Partial Authorizations,\" page 88.  - Split shipments: See \"Split Shipments,\" page 210. 
   * @member {String} linkId
   */
  exports.prototype['linkId'] = undefined;
  /**
   * Attribute that lets you define custom grouping for your processor reports. This field is supported only for **Worldpay VAP**.  See \"Report Groups,\" page 234. 
   * @member {String} reportGroup
   */
  exports.prototype['reportGroup'] = undefined;
  /**
   * Identifier for the **Visa Checkout** order. Visa Checkout provides a unique order ID for every transaction in the Visa Checkout **callID** field.  For more details, see Visa Checkout Using the SCMP API. 
   * @member {String} visaCheckoutId
   */
  exports.prototype['visaCheckoutId'] = undefined;
  /**
   * @member {module:model/Ptsv2paymentsIssuerInformation} issuer
   */
  exports.prototype['issuer'] = undefined;



  return exports;
}));


