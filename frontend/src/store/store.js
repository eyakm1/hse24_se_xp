import { configureStore } from '@reduxjs/toolkit';
import authReducer from './slices/authSlice';
import assignmentsReducer from './slices/assignmentsSlice';

export const store = configureStore({
  reducer: {
    auth: authReducer,
    assignments: assignmentsReducer,
  },
});

export default store;
