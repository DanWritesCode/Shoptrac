import { extend } from 'vee-validate'
import {
  required,
  // eslint-disable-next-line camelcase
  alpha_num,
  confirmed,
  max,
  min,
  email
} from 'vee-validate/dist/rules'

extend('required', {
  ...required,
  message: 'The {_field_} field is required.'
})
extend('alpha_num', {
  // eslint-disable-next-line camelcase
  ...alpha_num,
  message:
    'The {_field_} field must only contain Aa - zZ and 0 - 9  characters.'
})
extend('confirmed', {
  ...confirmed,
  message: 'The {_field_} field does not match.'
})
extend('max', {
  ...max,
  validate(value, { length }) {
    return value.length <= length
  },
  params: ['length'],
  message: 'The {_field_} field must have only {length} characters.'
})
extend('min', {
  ...min,
  validate(value, { length }) {
    return value.length >= length
  },
  params: ['length'],
  message: 'The {_field_} field must have at least {length} characters.'
})
extend('email', {
  ...email,
  message: 'The {_field_} is not a valid address.'
})
