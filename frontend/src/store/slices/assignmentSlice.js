import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import api from '../../api/api';

export const fetchAssignments = createAsyncThunk('assignments/fetchAssignments', async (_, { rejectWithValue }) => {
  try {
    const response = await api.get('/assignments');
    return response.data;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

const assignmentsSlice = createSlice({
  name: 'assignments',
  initialState: {
    items: [],
    loading: false,
    error: null,
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchAssignments.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchAssignments.fulfilled, (state, action) => {
        state.loading = false;
        state.items = action.payload;
      })
      .addCase(fetchAssignments.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export default assignmentsSlice.reducer;
