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

export const fetchAssignmentDetail = createAsyncThunk('assignments/fetchAssignmentDetail', async (id, { rejectWithValue }) => {
  try {
    const response = await api.get(`/assignments/${id}`);
    return response.data;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

export const submitAssignment = createAsyncThunk('assignments/submitAssignment', async (formData, { rejectWithValue }) => {
  try {
    const response = await api.post('/submissions', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

export const fetchSubmissionStatus = createAsyncThunk('assignments/fetchSubmissionStatus', async (assignmentId, { rejectWithValue }) => {
  try {
    const response = await api.get(`/submissions/${assignmentId}`);
    return response.data;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

export const fetchSubmissionsForAssignment = createAsyncThunk('assignments/fetchSubmissionsForAssignment', async (assignmentId, { rejectWithValue }) => {
  try {
    const response = await api.get(`/assignments/${assignmentId}/submissions`);
    return response.data;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

export const gradeSubmission = createAsyncThunk('assignments/gradeSubmission', async ({ submissionId, feedback, grade }, { rejectWithValue }) => {
  try {
    const response = await api.post(`/submissions/${submissionId}/grade`, { feedback, grade });
    return response.data;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

const assignmentsSlice = createSlice({
  name: 'assignments',
  initialState: {
    items: [],
    detail: null,
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
      })
      .addCase(fetchAssignmentDetail.pending, (state) => {
        state.loading = true;
        state.error = null;
        state.detail = null;
      })
      .addCase(fetchAssignmentDetail.fulfilled, (state, action) => {
        state.loading = false;
        state.detail = action.payload;
      })
      .addCase(fetchAssignmentDetail.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(submitAssignment.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(submitAssignment.fulfilled, (state) => {
        state.loading = false;
      })
      .addCase(submitAssignment.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(fetchSubmissionStatus.pending, (state) => {
        state.loading = true;
        state.error = null;
        state.submissionStatus = null;
      })
      .addCase(fetchSubmissionStatus.fulfilled, (state, action) => {
        state.loading = false;
        state.submissionStatus = action.payload;
      })
      .addCase(fetchSubmissionStatus.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      }).addCase(fetchSubmissionsForAssignment.pending, (state) => {
        state.loading = true;
        state.error = null;
        state.submissions = [];
      })
      .addCase(fetchSubmissionsForAssignment.fulfilled, (state, action) => {
        state.loading = false;
        state.submissions = action.payload;
      })
      .addCase(fetchSubmissionsForAssignment.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(gradeSubmission.pending, (state) => {
        state.loading = true;
      })
      .addCase(gradeSubmission.fulfilled, (state, action) => {
        state.loading = false;
        // Update submissions with new feedback and grade
        state.submissions = state.submissions.map((sub) =>
          sub.id === action.meta.arg.submissionId
            ? { ...sub, feedback: action.meta.arg.feedback, grade: action.meta.arg.grade }
            : sub
        );
      })
      .addCase(gradeSubmission.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export default assignmentsSlice.reducer;
