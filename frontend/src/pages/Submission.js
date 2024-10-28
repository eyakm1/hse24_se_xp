import React, { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { submitAssignment } from '../store/slices/assignmentsSlice';

function Submission() {
  const { id } = useParams();
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const { loading, error } = useSelector((state) => state.assignments);

  const [file, setFile] = useState(null);
  const [comment, setComment] = useState('');

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!file) {
      alert('Please select a file to submit.');
      return;
    }

    const formData = new FormData();
    formData.append('file', file);
    formData.append('comment', comment);
    formData.append('assignmentId', id);

    dispatch(submitAssignment(formData))
      .unwrap()
      .then(() => {
        alert('Submission successful!');
        navigate(`/assignments/${id}`);
      })
      .catch(() => alert('Submission failed. Please try again.'));
  };

  return (
    <div>
      <h1>Submit Assignment</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Upload File:
          <input type="file" onChange={handleFileChange} />
        </label>
        <label>
          Comment (optional):
          <textarea
            value={comment}
            onChange={(e) => setComment(e.target.value)}
            placeholder="Add any additional comments here"
          />
        </label>
        <button type="submit" disabled={loading}>
          {loading ? 'Submitting...' : 'Submit'}
        </button>
        {error && <p>Error: {error.message}</p>}
      </form>
    </div>
  );
}

export default Submission;
