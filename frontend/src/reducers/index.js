import { combineReducers } from 'redux'
import Auth from './user'
import books from "./books"
import error from "./error"
export default combineReducers({ Auth, books, error })