import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { login } from '../store/slices/authSlice';

function Login() {
  const dispatch = useDispatch();
  const { loading, error } = useSelector((state) => state.auth);

  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = (e) => {
    e.preventDefault();
    dispatch(login({ username, password }));
  };

  return (
    <form onSubmit={handleLogin}>
      <h1>Login</h1>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button type="submit" disabled={loading}>
        {loading ? 'Logging in...' : 'Login'}
      </button>
      {error && <p>{error.message || 'Login failed'}</p>}
    </form>
  );
}

export default Login;
