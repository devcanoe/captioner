import * as React from "react";
import { type VariantProps, cva } from "class-variance-authority";
import { cn } from "@/utils/lib/cn";

const ButtonVariants = cva("inline-flex item-center justify-center text-base leading-[24px] font-medium", {
  variants: {
    variant: {
      default: "bg-black p-2 rounded text-white hover:translate-y-[-1px] active:translate-y-[1px] transition-all duration-500",
      outline: "border rounded p-2",
      ghost: "hover:underline",
      icon: "",
    },
    size: {
      default: "w-20",
      sm: "",
      md: "",
      lg: "w-40",
    },
  },
  defaultVariants: {
    variant: "default",
    size: "default",
  },
});

interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof ButtonVariants> {}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, children, variant, size }, ref): React.JSX.Element => {
    return (
      <button
        className={cn(ButtonVariants({ size, variant, className }))}
        ref={ref}
      >
        {children}
      </button>
    );
  }
);

Button.displayName = "button";

export { Button, ButtonVariants };
