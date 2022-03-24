// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { GET_SESSION_ID_SUCCESS } from '@ttn-lw/lib/store/actions/init'

import {
  GET_USER,
  GET_USER_SUCCESS,
  GET_USER_FAILURE,
  UPDATE_USER_SUCCESS,
  LOGOUT_FAILURE,
  GET_USER_SESSIONS_LIST,
  GET_USER_SESSIONS_LIST_SUCCESS,
  GET_USER_SESSIONS_LIST_FAILURE,
} from '@account/store/actions/user'

const defaultState = {
  fetching: false,
  user: undefined,
  error: false,
  session_id: undefined,
}

const user = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_SESSION_ID_SUCCESS:
      return {
        ...state,
        session_id: payload,
      }
    case GET_USER:
      return {
        ...state,
        fetching: true,
        user: undefined,
        error: false,
      }
    case GET_USER_SUCCESS:
      return {
        ...state,
        user: payload,
        error: false,
      }
    case GET_USER_FAILURE:
      return {
        ...state,
        fetching: false,
        user: undefined,
        error: payload,
      }
    case UPDATE_USER_SUCCESS:
      return {
        ...state,
        user: {
          ...state.user,
          ...payload,
        },
      }
    case LOGOUT_FAILURE:
      return {
        ...state,
        fetching: false,
        error: payload,
      }
    case GET_USER_SESSIONS_LIST:
      return {
        ...state,
        fetching: true,
        user: state.user,
        error: false,
      }
    case GET_USER_SESSIONS_LIST_SUCCESS:
      return {
        ...state,
        user: {
          ...state.user,
          ...payload,
        },
        error: false,
      }
    case GET_USER_SESSIONS_LIST_FAILURE:
      return {
        ...state,
        user: state.user,
        fetching: false,
        error: payload,
      }
    default:
      return state
  }
}

export default user
