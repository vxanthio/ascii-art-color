# Team Permissions & Workflow

## Team Structure
- **Theo** (Leader): Repository Owner
- **Krystallenia**: Developer (Write permission)
- **Vasiliki**: Developer (Write permission)

---

## What EVERYONE Can Do (All Team Members)
✓ Create feature branches (feature/*, fix/*)
✓ Work on their own branches
✓ Commit and push to their feature branches
✓ Create Pull Requests
✓ Review each other's code
✓ Comment on PRs and issues
✓ Clone and fork the repository

---

## What ONLY THEO Can Do (Owner Permission)
✓ Merge Pull Requests to `develop` and `main`
✓ Push directly to `develop` and `main` (protected branches)
✓ Delete branches
✓ Change repository settings
✓ Manage collaborators and permissions
✓ Protect/unprotect branches
✓ Configure webhooks and integrations

---

## What Krystallenia & Vasiliki CANNOT Do (Write Permission Restrictions)
✗ Push directly to `main` or `develop` (protected branches)
✗ Merge Pull Requests (even their own)
✗ Delete protected branches
✗ Change repository settings
✗ Add/remove collaborators
✗ Modify branch protection rules

**Why these restrictions?**
→ Ensures all code goes through review process
→ Protects main/develop from accidental breaks
→ Maintains code quality standards
→ Industry-standard team workflow

---

## Standard Workflow

**For Krystallenia & Vasiliki:**
```bash
1. git checkout develop
2. git pull origin develop
3. git checkout -b feature/your-feature
4. [make changes and commit]
5. git push origin feature/your-feature
6. Create PR on Gitea: feature/your-feature → develop
7. Wait for Theo's review
8. Address feedback if needed
9. Theo merges when approved
```

**For Theo:**
```bash
Same as above, PLUS:
- Review and merge PRs from team
- Can push directly to develop if needed (avoid this!)
- Merge develop → main for releases
```

---

## PR Review Process

**Theo reviews PRs for:**
□ All tests passing
□ Follows conventional commit format
□ Code quality and readability
□ Proper error handling
□ No merge conflicts
□ Jira ticket linked

---

## Questions or Issues?

Contact Theo immediately if:
- You're blocked by permissions
- Something seems wrong with access
- Workflow is unclear
- You need to do something urgent

---

## Remember
This permission structure is intentional and follows industry best practices. It's not about trust—it's about maintaining code quality through mandatory review processes. Even Theo should create PRs for major features! 👍
```

---

## 📊 CLEARER VISUALIZATION
```
╔══════════════════════════════════════════════════════════╗
║                    PERMISSION MATRIX                      ║
╠══════════════════════════════════════════════════════════╣
║ Action                        │ Theo │ Kryst. │ Vasil.  ║
╠═══════════════════════════════╪══════╪════════╪═════════╣
║ Create feature branches       │  ✓   │   ✓    │   ✓     ║
║ Push to feature branches      │  ✓   │   ✓    │   ✓     ║
║ Create PRs                    │  ✓   │   ✓    │   ✓     ║
║ Review code                   │  ✓   │   ✓    │   ✓     ║
║ Comment on PRs                │  ✓   │   ✓    │   ✓     ║
║───────────────────────────────┼──────┼────────┼─────────║
║ Push to develop               │  ✓   │   ✗    │   ✗     ║
║ Push to main                  │  ✓   │   ✗    │   ✗     ║
║ Merge PRs                     │  ✓   │   ✗    │   ✗     ║
║ Delete branches               │  ✓   │   ✗    │   ✗     ║
║ Change settings               │  ✓   │   ✗    │   ✗     ║
║ Manage permissions            │  ✓   │   ✗    │   ✗     ║
╚══════════════════════════════════════════════════════════╝