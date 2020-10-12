import { extend } from 'vee-validate';
import {
  required,
  email,
  numeric,
  regex
} from 'vee-validate/dist/rules'

extend('required', {
  ...required,
  message: 'This field is required',
})
extend('email', {
  ...email,
  message: '{_field_} must be a valid email',
})