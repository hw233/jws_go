define(['exports', 'module', 'react', 'classnames', './BootstrapMixin', './utils/ValidComponentChildren', './Glyphicon'], function (exports, module, _react, _classnames, _BootstrapMixin, _utilsValidComponentChildren, _Glyphicon) {
  'use strict';

  var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

  function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

  var _React = _interopRequireDefault(_react);

  var _classNames = _interopRequireDefault(_classnames);

  var _BootstrapMixin2 = _interopRequireDefault(_BootstrapMixin);

  var _ValidComponentChildren = _interopRequireDefault(_utilsValidComponentChildren);

  var _Glyphicon2 = _interopRequireDefault(_Glyphicon);

  var Carousel = _React['default'].createClass({
    displayName: 'Carousel',

    mixins: [_BootstrapMixin2['default']],

    propTypes: {
      slide: _React['default'].PropTypes.bool,
      indicators: _React['default'].PropTypes.bool,
      interval: _React['default'].PropTypes.number,
      controls: _React['default'].PropTypes.bool,
      pauseOnHover: _React['default'].PropTypes.bool,
      wrap: _React['default'].PropTypes.bool,
      onSelect: _React['default'].PropTypes.func,
      onSlideEnd: _React['default'].PropTypes.func,
      activeIndex: _React['default'].PropTypes.number,
      defaultActiveIndex: _React['default'].PropTypes.number,
      direction: _React['default'].PropTypes.oneOf(['prev', 'next']),
      prevIcon: _React['default'].PropTypes.node,
      nextIcon: _React['default'].PropTypes.node
    },

    getDefaultProps: function getDefaultProps() {
      return {
        slide: true,
        interval: 5000,
        pauseOnHover: true,
        wrap: true,
        indicators: true,
        controls: true,
        prevIcon: _React['default'].createElement(_Glyphicon2['default'], { glyph: 'chevron-left' }),
        nextIcon: _React['default'].createElement(_Glyphicon2['default'], { glyph: 'chevron-right' })
      };
    },

    getInitialState: function getInitialState() {
      return {
        activeIndex: this.props.defaultActiveIndex == null ? 0 : this.props.defaultActiveIndex,
        previousActiveIndex: null,
        direction: null
      };
    },

    getDirection: function getDirection(prevIndex, index) {
      if (prevIndex === index) {
        return null;
      }

      return prevIndex > index ? 'prev' : 'next';
    },

    componentWillReceiveProps: function componentWillReceiveProps(nextProps) {
      var activeIndex = this.getActiveIndex();

      if (nextProps.activeIndex != null && nextProps.activeIndex !== activeIndex) {
        clearTimeout(this.timeout);
        this.setState({
          previousActiveIndex: activeIndex,
          direction: nextProps.direction != null ? nextProps.direction : this.getDirection(activeIndex, nextProps.activeIndex)
        });
      }
    },

    componentDidMount: function componentDidMount() {
      this.waitForNext();
    },

    componentWillUnmount: function componentWillUnmount() {
      clearTimeout(this.timeout);
    },

    next: function next(e) {
      if (e) {
        e.preventDefault();
      }

      var index = this.getActiveIndex() + 1;
      var count = _ValidComponentChildren['default'].numberOf(this.props.children);

      if (index > count - 1) {
        if (!this.props.wrap) {
          return;
        }
        index = 0;
      }

      this.handleSelect(index, 'next');
    },

    prev: function prev(e) {
      if (e) {
        e.preventDefault();
      }

      var index = this.getActiveIndex() - 1;

      if (index < 0) {
        if (!this.props.wrap) {
          return;
        }
        index = _ValidComponentChildren['default'].numberOf(this.props.children) - 1;
      }

      this.handleSelect(index, 'prev');
    },

    pause: function pause() {
      this.isPaused = true;
      clearTimeout(this.timeout);
    },

    play: function play() {
      this.isPaused = false;
      this.waitForNext();
    },

    waitForNext: function waitForNext() {
      if (!this.isPaused && this.props.slide && this.props.interval && this.props.activeIndex == null) {
        this.timeout = setTimeout(this.next, this.props.interval);
      }
    },

    handleMouseOver: function handleMouseOver() {
      if (this.props.pauseOnHover) {
        this.pause();
      }
    },

    handleMouseOut: function handleMouseOut() {
      if (this.isPaused) {
        this.play();
      }
    },

    render: function render() {
      var classes = {
        carousel: true,
        slide: this.props.slide
      };

      return _React['default'].createElement(
        'div',
        _extends({}, this.props, {
          className: (0, _classNames['default'])(this.props.className, classes),
          onMouseOver: this.handleMouseOver,
          onMouseOut: this.handleMouseOut }),
        this.props.indicators ? this.renderIndicators() : null,
        _React['default'].createElement(
          'div',
          { className: 'carousel-inner', ref: 'inner' },
          _ValidComponentChildren['default'].map(this.props.children, this.renderItem)
        ),
        this.props.controls ? this.renderControls() : null
      );
    },

    renderPrev: function renderPrev() {
      return _React['default'].createElement(
        'a',
        { className: 'left carousel-control', href: '#prev', key: 0, onClick: this.prev },
        this.props.prevIcon
      );
    },

    renderNext: function renderNext() {
      return _React['default'].createElement(
        'a',
        { className: 'right carousel-control', href: '#next', key: 1, onClick: this.next },
        this.props.nextIcon
      );
    },

    renderControls: function renderControls() {
      if (!this.props.wrap) {
        var activeIndex = this.getActiveIndex();
        var count = _ValidComponentChildren['default'].numberOf(this.props.children);

        return [activeIndex !== 0 ? this.renderPrev() : null, activeIndex !== count - 1 ? this.renderNext() : null];
      }

      return [this.renderPrev(), this.renderNext()];
    },

    renderIndicator: function renderIndicator(child, index) {
      var className = index === this.getActiveIndex() ? 'active' : null;

      return _React['default'].createElement('li', {
        key: index,
        className: className,
        onClick: this.handleSelect.bind(this, index, null) });
    },

    renderIndicators: function renderIndicators() {
      var indicators = [];
      _ValidComponentChildren['default'].forEach(this.props.children, function (child, index) {
        indicators.push(this.renderIndicator(child, index),

        // Force whitespace between indicator elements, bootstrap
        // requires this for correct spacing of elements.
        ' ');
      }, this);

      return _React['default'].createElement(
        'ol',
        { className: 'carousel-indicators' },
        indicators
      );
    },

    getActiveIndex: function getActiveIndex() {
      return this.props.activeIndex != null ? this.props.activeIndex : this.state.activeIndex;
    },

    handleItemAnimateOutEnd: function handleItemAnimateOutEnd() {
      this.setState({
        previousActiveIndex: null,
        direction: null
      }, function () {
        this.waitForNext();

        if (this.props.onSlideEnd) {
          this.props.onSlideEnd();
        }
      });
    },

    renderItem: function renderItem(child, index) {
      var activeIndex = this.getActiveIndex();
      var isActive = index === activeIndex;
      var isPreviousActive = this.state.previousActiveIndex != null && this.state.previousActiveIndex === index && this.props.slide;

      return (0, _react.cloneElement)(child, {
        active: isActive,
        ref: child.ref,
        key: child.key ? child.key : index,
        index: index,
        animateOut: isPreviousActive,
        animateIn: isActive && this.state.previousActiveIndex != null && this.props.slide,
        direction: this.state.direction,
        onAnimateOutEnd: isPreviousActive ? this.handleItemAnimateOutEnd : null
      });
    },

    handleSelect: function handleSelect(index, direction) {
      clearTimeout(this.timeout);

      var previousActiveIndex = this.getActiveIndex();
      direction = direction || this.getDirection(previousActiveIndex, index);

      if (this.props.onSelect) {
        this.props.onSelect(index, direction);
      }

      if (this.props.activeIndex == null && index !== previousActiveIndex) {
        if (this.state.previousActiveIndex != null) {
          // If currently animating don't activate the new index.
          // TODO: look into queuing this canceled call and
          // animating after the current animation has ended.
          return;
        }

        this.setState({
          activeIndex: index,
          previousActiveIndex: previousActiveIndex,
          direction: direction
        });
      }
    }
  });

  module.exports = Carousel;
});