import { extend } from 'vee-validate';
import {
  required,
  email,
  numeric,
  regex,
  alpha
} from 'vee-validate/dist/rules'

extend('required', {
  ...required,
  message: '{_field_} is required',
})
extend('email', {
  ...email,
  message: '{_field_} must be a valid email',
})
extend('alpha', {
  ...alpha,
  message: '{_field_} must only contain letters'
})