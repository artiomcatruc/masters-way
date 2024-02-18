import {useState} from "react";
import {RouterProvider} from "react-router-dom";
import {
  DEFAULT_NOTIFICATION_SETTINGS,
  globalContext,
} from "src/GlobalContext";
import {UserPreview} from "src/model/businessModelPreview/UserPreview";
import {router} from "src/router/Router";

/**
 * App
 */
export const App = () => {
  const [user, setUser] = useState<UserPreview | null>(null);
  const [isInitialized, setIsInitialized] = useState(false);

  return (
    <globalContext.Provider value={{
      user,
      setUser,
      isInitialized,
      setIsInitialized,
      // TODO: load from local storage
      notification: DEFAULT_NOTIFICATION_SETTINGS,
    }}
    >
      <RouterProvider router={router} />
    </globalContext.Provider>
  );
};
