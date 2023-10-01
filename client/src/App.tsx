import { AuthProvider } from './auth/auth-context';
import PrivateRoute from './auth/private-route';
import Toolbar from './components/toolbar';
import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <AuthProvider>
        <Router>
          <Toolbar />
          <Routes>
            <Route path="/" element={<PrivateRoute />}>
              <Route path="/authroute" element={<div>AuthRoute</div>} />
            </Route>
          </Routes>
        </Router>
      </AuthProvider>
    </div>
  );
}

export default App;
