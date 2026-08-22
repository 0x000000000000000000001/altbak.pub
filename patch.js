const fs = require('fs');
const file = '/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs';
let content = fs.readFileSync(file, 'utf8');

content = content.replace(
`          genCurried n args = 
            let argName = "a" <> show (len - n)
                argTy = fromMaybe Any (Array.index argTys (len - n))
                argTyStr = codegenExprType currentMod false argTy
                nextRetTyStr = if n == 1 then retTyStr else "std::rc::Rc<dyn Fn(" <> codegenExprType currentMod false (fromMaybe Any (Array.index argTys (len - n + 1))) <> ") -> " <> codegenExprType currentMod true (if n > 2 then Func (Array.drop (len - n + 2) argTys) retTy else if n == 2 then retTy else Any) <> ">"
            in "std::rc::Rc::new(move |mut " <> argName <> ": " <> argTyStr <> "| -> " <> nextRetTyStr <> " { " <> genCurried (n - 1) (Array.snoc args argName) <> " })"`,
`          genCurried n args = 
            let argName = "a" <> show (len - n)
                argTy = fromMaybe Any (Array.index argTys (len - n))
                argTyStr = codegenExprType currentMod false argTy
                nextRetTyStr = if n == 1 then retTyStr else "std::rc::Rc<dyn Fn(" <> codegenExprType currentMod false (fromMaybe Any (Array.index argTys (len - n + 1))) <> ") -> " <> codegenExprType currentMod true (if n > 2 then Func (Array.drop (len - n + 2) argTys) retTy else if n == 2 then retTy else Any) <> ">"
                clonesCode = if n > 1 && Array.length args > 0 then String.joinWith " " (map (\\a -> "let mut " <> a <> " = " <> a <> ".clone();") args) <> " " else ""
            in "std::rc::Rc::new(move |mut " <> argName <> ": " <> argTyStr <> "| -> " <> nextRetTyStr <> " { " <> clonesCode <> genCurried (n - 1) (Array.snoc args argName) <> " })"`);

fs.writeFileSync(file, content);
