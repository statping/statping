/// <reference types="cypress" />

import "../support/commands"

context('Users Tests', () => {


  beforeEach(() => {
    cy.restoreLocalStorageCache();
  });

  afterEach(() => {
    cy.saveLocalStorageCache();
  });

  it('should Login', () => {
    cy.visit('/login')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('.navbar-brand').should('contain', 'Statping')
    cy.getCookies()

    cy.getCookies().should('have.length', 1)
  })

  it('should goto users', () => {
    cy.visit('/dashboard/users')
    cy.get('#users_table > tr').should('have.length', 1)
    cy.get('#users_table > tr').eq(0).contains('admin')
  })

  it('should create new User', () => {
    cy.visit('/dashboard/users')
    cy.get('#username').clear().type('admin2')
    cy.get('#email').clear().type('info@admin.com')
    cy.get('#password').clear().type('password123')
    cy.get('#password_confirm').clear().type('password123')

    cy.get('button[type="submit"]').click()
    cy.get('#users_table > tr').should('have.length', 2)
  })

  it('should create new Admin User', () => {
    cy.visit('/dashboard/users')
    cy.get('#username').clear().type('admin3')
    cy.get('#admin_switch').click()
    cy.get('#email').clear().type('info@admin3.com')
    cy.get('#password').clear().type('password123')
    cy.get('#password_confirm').clear().type('password123')

    cy.get('button[type="submit"]').click()
    cy.get('#users_table > tr').should('have.length', 3)
  })

  it('should confirm new user', () => {
    cy.visit('/dashboard/users')
    cy.get('#users_table > tr').should('have.length', 3)
    cy.get('#users_table > tr').eq(0).contains('admin')
    cy.get('#users_table > tr').eq(1).contains('admin2')
    cy.get('#users_table > tr').eq(2).contains('admin3')

    cy.get('#users_table > tr').eq(0).contains('ADMIN')
    cy.get('#users_table > tr').eq(1).contains('USER')
    cy.get('#users_table > tr').eq(2).contains('ADMIN')
  })

  it('should delete new users', () => {
    cy.visit('/dashboard/users')
    cy.get('#users_table > tr').should('have.length', 3)
    cy.get('#users_table > tr').eq(2).find('a.btn-danger').click()
    cy.get('#users_table > tr').eq(1).find('a.btn-danger').click()
    cy.get('#users_table > tr').should('have.length', 1)
  })

})
