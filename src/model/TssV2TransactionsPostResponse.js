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
    root.CyberSource.TssV2TransactionsPostResponse = factory(root.CyberSource.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The TssV2TransactionsPostResponse model module.
   * @module model/TssV2TransactionsPostResponse
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>TssV2TransactionsPostResponse</code>.
   * @alias module:model/TssV2TransactionsPostResponse
   * @class
   */
  var exports = function() {
    var _this = this;








  };

  /**
   * Constructs a <code>TssV2TransactionsPostResponse</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TssV2TransactionsPostResponse} obj Optional instance to populate.
   * @return {module:model/TssV2TransactionsPostResponse} The populated <code>TssV2TransactionsPostResponse</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('save')) {
        obj['save'] = ApiClient.convertToType(data['save'], 'Boolean');
      }
      if (data.hasOwnProperty('name')) {
        obj['name'] = ApiClient.convertToType(data['name'], 'String');
      }
      if (data.hasOwnProperty('timezone')) {
        obj['timezone'] = ApiClient.convertToType(data['timezone'], 'String');
      }
      if (data.hasOwnProperty('query')) {
        obj['query'] = ApiClient.convertToType(data['query'], 'String');
      }
      if (data.hasOwnProperty('offset')) {
        obj['offset'] = ApiClient.convertToType(data['offset'], 'Number');
      }
      if (data.hasOwnProperty('limit')) {
        obj['limit'] = ApiClient.convertToType(data['limit'], 'Number');
      }
      if (data.hasOwnProperty('sort')) {
        obj['sort'] = ApiClient.convertToType(data['sort'], 'String');
      }
    }
    return obj;
  }

  /**
   * save or not save.
   * @member {Boolean} save
   */
  exports.prototype['save'] = undefined;
  /**
   * The description for this field is not available. 
   * @member {String} name
   */
  exports.prototype['name'] = undefined;
  /**
   * Time Zone.
   * @member {String} timezone
   */
  exports.prototype['timezone'] = undefined;
  /**
   * transaction search query string.
   * @member {String} query
   */
  exports.prototype['query'] = undefined;
  /**
   * offset.
   * @member {Number} offset
   */
  exports.prototype['offset'] = undefined;
  /**
   * limit on number of results.
   * @member {Number} limit
   */
  exports.prototype['limit'] = undefined;
  /**
   * A comma separated list of the following form - fieldName1 asc or desc, fieldName2 asc or desc, etc.
   * @member {String} sort
   */
  exports.prototype['sort'] = undefined;



  return exports;
}));


