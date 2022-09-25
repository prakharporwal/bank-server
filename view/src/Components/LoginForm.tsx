import {
  FormControl,
  FormLabel,
  FormErrorMessage,
  FormHelperText,
  Input,
  Button,
  useToast,
} from "@chakra-ui/react";
import { time } from "console";
import { EventHandler, useEffect, useState } from "react";
import { utils } from "../utils/cookie";

var cronId: NodeJS.Timeout;

// TODO: validate form values
const LoginForm: React.FunctionComponent<any> = (props) => {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const toast = useToast();

  let isError = email === "";

  const handleEmailChange = (e: any) => {
    setEmail(e.target.value);
  };
  const handlePasswordChange = (e: any) => setPassword(e.target.value);

  const handleSubmit = async (e: any) => {
    e.preventDefault();

    setIsLoading(true);

    let reqbody: any = { user_id: email, password: password };

    fetch("http://localhost:8080/public/v1/login", {
      method: "POST",
      body: JSON.stringify(reqbody),
    })
      .then((res) => {
        if (!res.ok) {
          let toastId = "login";
          if (!toast.isActive(toastId)) {
            toast({
              id: toastId,
              title: "LogIn Failed",
              description: "for user : " + reqbody.user_id,
              status: "error",
              duration: 4000,
              isClosable: true,
            });
          }

          throw new Error("Login Failed");
        }

        return res.json();
      })
      .then((res) => {
        signOut(); // to clear setInterval loop

        // set token in cookie
        document.cookie = "access_token=".concat(res.access_token);
        document.cookie = "refresh_token=".concat(res.refresh_token);
        document.cookie = "expires_at=".concat(res.expires_at);

        let date = new Date(res.expires_at);

        let toastId = "login";
        if (!toast.isActive(toastId)) {
          toast({
            id: toastId,
            title: "Logged In",
            description: "User : " + reqbody.user_id,
            status: "success",
            duration: 4000,
            isClosable: true,
          });
        }
      })
      .catch((err) => console.error(err))
      .finally(() => {
        setIsLoading(false);
      });
  };

  useEffect(() => {
    refreshTokenCron();
  }, []);

  return (
    <form onSubmit={handleSubmit}>
      <FormControl isInvalid={isError}>
        <FormLabel>Email</FormLabel>
        <Input
          type="email"
          value={email}
          onChange={handleEmailChange}
          isRequired
        />
        {!isError ? (
          <FormHelperText>
            Enter the email you'd like to receive the newsletter on.
          </FormHelperText>
        ) : (
          <FormErrorMessage>Email is required.</FormErrorMessage>
        )}
      </FormControl>
      <FormControl>
        <FormLabel>Password</FormLabel>
        <Input
          type="password"
          value={password}
          onChange={handlePasswordChange}
          isRequired
        />
      </FormControl>

      <Button
        isLoading={isLoading}
        type="submit"
        variant="outline"
        width="full"
        mt={4}
        onSubmit={(e) => handleSubmit(e)}
      >
        Sign In
      </Button>
    </form>
  );
};

// TODO: validate form values
const SignUpForm: React.FunctionComponent<any> = (props) => {
  const [email, setEmail] = useState<string>("");
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const [isLoading, setIsLoading] = useState<boolean>(false);

  const toast = useToast();

  let isError = email === "";

  const handleEmailChange = (e: any) => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = (e: any) => setPassword(e.target.value);
  const handleUsernameChange = (e: any) => setUsername(e.target.value);

  const handleSubmit = async (e: any) => {
    e.preventDefault();

    setIsLoading(true);

    console.log("hello");
    let reqbody: any = { email: email, password: password, username: username };
    console.log(utils.getCookie("access_token"));

    fetch("http://localhost:8080/public/v1/signup", {
      method: "POST",
      body: JSON.stringify(reqbody),
      headers: new Headers({
        access_token: utils.getCookie("access_token"),
      }),
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res);

        let toastId = "signup";
        if (!toast.isActive(toastId)) {
          toast({
            id: toastId,
            title: "Account created",
            description: "We've created your account for you.",
            status: "success",
            duration: 4000,
            isClosable: true,
          });
        }
      })
      .catch((err) => console.error(err))
      .finally(() => {
        setIsLoading(false);
      });
  };

  return (
    <form onSubmit={handleSubmit}>
      <FormControl isInvalid={isError}>
        <FormLabel>Username</FormLabel>
        <Input
          type="username"
          value={username}
          onChange={handleUsernameChange}
          isRequired
        />
        {!isError ? (
          <FormHelperText>
            Enter the email you'd like to receive the newsletter on.
          </FormHelperText>
        ) : (
          <FormErrorMessage>Email is required.</FormErrorMessage>
        )}
      </FormControl>
      <FormControl isInvalid={isError}>
        <FormLabel>Email</FormLabel>
        <Input
          type="email"
          value={email}
          onChange={handleEmailChange}
          isRequired
        />
        {!isError ? (
          <FormHelperText>
            Enter the email you'd like to receive the newsletter on.
          </FormHelperText>
        ) : (
          <FormErrorMessage>Email is required.</FormErrorMessage>
        )}
      </FormControl>
      <FormControl>
        <FormLabel>Password</FormLabel>
        <Input
          type="password"
          value={password}
          onChange={handlePasswordChange}
          isRequired
        />
      </FormControl>

      <Button
        isLoading={isLoading}
        type="submit"
        variant="outline"
        width="full"
        mt={4}
      >
        Sign Up
      </Button>
    </form>
  );
};

function refreshTokenCron() {
  console.log("triggered refresh token cron!");

  // refresh cron rate
  const FourMinutes: number = 240000;

  cronId = setInterval(() => {
    let date = new Date(utils.getCookie("expires_at"));

    if (date.getDate() - Date.now() > 60000) {
      return;
    }

    let reqbody = {
      refresh_token: utils.getCookie("refresh_token"),
    };

    fetch("http://localhost:8080/v1/session/refresh", {
      method: "POST",
      body: JSON.stringify(reqbody),
      headers: new Headers({
        access_token: utils.getCookie("access_token"),
      }),
    })
      .then((res) => {
        if (!res.ok) {
          signOut();
          clearTimeout(cronId);

          throw new Error("failed refreshing session!");
        }
        return res.json();
      })
      .then((res) => {
        document.cookie = "access_token=".concat(res.access_token);
        document.cookie = "refresh_token=".concat(res.refresh_token);
        document.cookie = "expires_at=".concat(res.expires_at);

        let date = new Date(res.expires_at);
        console.log("session refreshed ", res.refresh_token);
      })
      .catch((err) => console.error(err))
      .finally(() => {});
  }, FourMinutes);
}

function signOut() {
  clearTimeout(cronId);
  document.cookie = "access_token=";
  document.cookie = "refresh_token=";
  document.cookie = "expires_at=";
}

export { LoginForm, SignUpForm };
