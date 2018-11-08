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
    define(['ApiClient', 'model/ReportingV3ReportSubscriptionsGet200ResponseSubscriptions'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./ReportingV3ReportSubscriptionsGet200ResponseSubscriptions'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.ReportingV3ReportSubscriptionsGet200Response = factory(root.CyberSource.ApiClient, root.CyberSource.ReportingV3ReportSubscriptionsGet200ResponseSubscriptions);
  }
}(this, function(ApiClient, ReportingV3ReportSubscriptionsGet200ResponseSubscriptions) {
  'use strict';




  /**
   * The ReportingV3ReportSubscriptionsGet200Response model module.
   * @module model/ReportingV3ReportSubscriptionsGet200Response
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>ReportingV3ReportSubscriptionsGet200Response</code>.
   * @alias module:model/ReportingV3ReportSubscriptionsGet200Response
   * @class
   */
  var exports = function() {
    var _this = this;


  };

  /**
   * Constructs a <code>ReportingV3ReportSubscriptionsGet200Response</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ReportingV3ReportSubscriptionsGet200Response} obj Optional instance to populate.
   * @return {module:model/ReportingV3ReportSubscriptionsGet200Response} The populated <code>ReportingV3ReportSubscriptionsGet200Response</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('subscriptions')) {
        obj['subscriptions'] = ApiClient.convertToType(data['subscriptions'], [ReportingV3ReportSubscriptionsGet200ResponseSubscriptions]);
      }
    }
    return obj;
  }

  /**
   * @member {Array.<module:model/ReportingV3ReportSubscriptionsGet200ResponseSubscriptions>} subscriptions
   */
  exports.prototype['subscriptions'] = undefined;



  return exports;
}));


