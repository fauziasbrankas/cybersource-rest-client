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
    define(['ApiClient', 'model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200ResponseLinks'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./TmsV1InstrumentidentifiersPaymentinstrumentsGet200ResponseLinks'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response = factory(root.CyberSource.ApiClient, root.CyberSource.TmsV1InstrumentidentifiersPaymentinstrumentsGet200ResponseLinks);
  }
}(this, function(ApiClient, TmsV1InstrumentidentifiersPaymentinstrumentsGet200ResponseLinks) {
  'use strict';




  /**
   * The TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response model module.
   * @module model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response</code>.
   * @alias module:model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response
   * @class
   */
  var exports = function() {
    var _this = this;








  };

  /**
   * Constructs a <code>TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response} obj Optional instance to populate.
   * @return {module:model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response} The populated <code>TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('_links')) {
        obj['_links'] = TmsV1InstrumentidentifiersPaymentinstrumentsGet200ResponseLinks.constructFromObject(data['_links']);
      }
      if (data.hasOwnProperty('object')) {
        obj['object'] = ApiClient.convertToType(data['object'], 'String');
      }
      if (data.hasOwnProperty('offset')) {
        obj['offset'] = ApiClient.convertToType(data['offset'], 'String');
      }
      if (data.hasOwnProperty('limit')) {
        obj['limit'] = ApiClient.convertToType(data['limit'], 'String');
      }
      if (data.hasOwnProperty('count')) {
        obj['count'] = ApiClient.convertToType(data['count'], 'String');
      }
      if (data.hasOwnProperty('total')) {
        obj['total'] = ApiClient.convertToType(data['total'], 'String');
      }
      if (data.hasOwnProperty('_embedded')) {
        obj['_embedded'] = ApiClient.convertToType(data['_embedded'], Object);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200ResponseLinks} _links
   */
  exports.prototype['_links'] = undefined;
  /**
   * Shows the response is a collection of objects.
   * @member {module:model/TmsV1InstrumentidentifiersPaymentinstrumentsGet200Response.ObjectEnum} object
   */
  exports.prototype['object'] = undefined;
  /**
   * The offset parameter supplied in the request.
   * @member {String} offset
   */
  exports.prototype['offset'] = undefined;
  /**
   * The limit parameter supplied in the request.
   * @member {String} limit
   */
  exports.prototype['limit'] = undefined;
  /**
   * The number of Payment Instruments returned in the array.
   * @member {String} count
   */
  exports.prototype['count'] = undefined;
  /**
   * The total number of Payment Instruments associated with the Instrument Identifier in the zero-based dataset.
   * @member {String} total
   */
  exports.prototype['total'] = undefined;
  /**
   * Array of Payment Instruments returned for the supplied Instrument Identifier.
   * @member {Object} _embedded
   */
  exports.prototype['_embedded'] = undefined;


  /**
   * Allowed values for the <code>object</code> property.
   * @enum {String}
   * @readonly
   */
  exports.ObjectEnum = {
    /**
     * value: "collection"
     * @const
     */
    "collection": "collection"  };


  return exports;
}));


